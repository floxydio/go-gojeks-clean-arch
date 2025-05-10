package clientorder

import (
	"context"
	"encoding/json"
	"fmt"
	"gojeksrepo/ent"
	"gojeksrepo/ent/trip"
	"gojeksrepo/internal/client/client_order/dto"
	"gojeksrepo/pkg"
	"time"

	"github.com/google/uuid"
	"github.com/segmentio/kafka-go"
)

type ClientOrderService struct {
	dbClient     *ent.Client
	kafkaService *kafka.Conn
}

func NewServiceClientOrderService(db *ent.Client, kafka *kafka.Conn) *ClientOrderService {
	return &ClientOrderService{
		dbClient:     db,
		kafkaService: kafka,
	}
}

func (repo *ClientOrderService) FindLastTransactionNotSuccess(ctx context.Context) ([]*ent.Trip, error) {
	dataTrx, err := repo.dbClient.Trip.Query().Where(trip.StatusIn(trip.StatusCancelled)).All(ctx)

	if err != nil {
		return nil, err
	}

	return dataTrx, nil
}

func (repo *ClientOrderService) CreateOrder(formData dto.TripRequest, ctx context.Context) (*ent.Trip, error) {

	fmt.Println(formData.UserID)
	uuidUser, errUUID := uuid.Parse(formData.UserID)

	if errUUID != nil {
		return nil, errUUID
	}

	data, errSend := repo.dbClient.Trip.Create().SetPickupLat(formData.PickupLat).SetPickupLong(formData.PickupLong).SetDropLat(formData.DropLat).SetDropLong(formData.DropLong).SetStatus(trip.StatusRequested).SetDistanceKm(formData.DistanceKm).SetUserID(uuidUser).SetIsPaid(formData.IsPaid).Save(ctx)

	if errSend != nil {
		return nil, errSend
	}

	var timeGenerate = time.Now()

	var modelToKafka struct {
		OrderId      uuid.UUID   `json:"order_id"`
		UserId       uuid.UUID   `json:"user_id"`
		StatusOrder  trip.Status `json:"status_order"`
		TimeGenerate string      `json:"time_generate"`
	}

	modelToKafka.OrderId = data.ID
	modelToKafka.UserId = data.UserID
	modelToKafka.StatusOrder = trip.StatusRequested
	modelToKafka.TimeGenerate = timeGenerate.Format("2006-01-02")

	jsonKafkaModel, _ := json.Marshal(modelToKafka)

	_, errKafka := repo.kafkaService.WriteMessages(
		kafka.Message{Key: []byte(modelToKafka.OrderId.String()), Value: []byte((jsonKafkaModel))},
	)

	if errKafka != nil {
		return nil, errKafka
	}

	pkg.SendToUser(formData.UserID, "Order berhasil dibuat. Menunggu driver...")

	return data, nil

}
