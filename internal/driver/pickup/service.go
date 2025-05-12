package pickup

import (
	"context"
	"encoding/json"
	"fmt"
	"gojeksrepo/ent"
	"gojeksrepo/ent/trip"
	"gojeksrepo/ent/user"
	"gojeksrepo/pkg"

	"github.com/google/uuid"
	"github.com/segmentio/kafka-go"
)

type PickupServiceClient struct {
	dbEnt             *ent.Client
	kafkaClientDriver *kafka.Writer
}

func PickupNewService(db *ent.Client, kafka *kafka.Writer) *PickupServiceClient {
	return &PickupServiceClient{
		dbEnt:             db,
		kafkaClientDriver: kafka,
	}
}

func (svc *PickupServiceClient) AcceptOrder(orderId string, driverId string, ctx context.Context) error {
	findUser, errFindFirst := svc.dbEnt.Trip.Query().Where(trip.ID(uuid.MustParse(orderId))).First(ctx)

	if errFindFirst != nil {
		return errFindFirst
	}

	findDriverId, errFindDriver := svc.dbEnt.User.Query().Where(user.ID(uuid.MustParse(driverId))).First(ctx)

	if errFindDriver != nil {
		return errFindDriver
	}
	_, err := svc.dbEnt.Trip.Update().Where(trip.ID(uuid.MustParse(orderId))).SetDriverID(uuid.MustParse(driverId)).SetStatus(trip.StatusAccepted).Save(ctx)

	if err != nil {
		return err
	}

	var jsonModel struct {
		OrderId  uuid.UUID   `json:"order_id"`
		DriverId uuid.UUID   `json:"driver_id"`
		Status   trip.Status `json:"status"`
	}

	jsonModel.DriverId = uuid.MustParse(driverId)
	jsonModel.OrderId = uuid.MustParse(orderId)
	jsonModel.Status = trip.StatusAccepted

	jsonData, err := json.Marshal(jsonModel)
	if err != nil {
		return fmt.Errorf("failed to marshal kafka message: %w", err)
	}

	err = svc.kafkaClientDriver.WriteMessages(ctx, kafka.Message{
		Key:   []byte(orderId),
		Value: []byte(jsonData),
	})

	if err != nil {
		return err
	}

	pkg.SendToUser(findUser.UserID.String(), fmt.Sprintf("Orderan anda telah dipickup oleh %s", findDriverId.Name))

	return err
}
