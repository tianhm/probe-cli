package mocks

import (
	"context"

	"github.com/ooni/probe-cli/v3/internal/model"
)

// Experiment mocks model.Experiment
type Experiment struct {
	MockKibiBytesReceived func() float64

	MockKibiBytesSent func() float64

	MockName func() string

	MockReportID func() string

	MockMeasureWithContext func(
		ctx context.Context, target model.ExperimentTarget) (measurement *model.Measurement, err error)

	MockSaveMeasurement func(measurement *model.Measurement, filePath string) error

	MockSubmitAndUpdateMeasurementContext func(
		ctx context.Context, measurement *model.Measurement) (string, error)

	MockOpenReportContext func(ctx context.Context) error
}

func (e *Experiment) KibiBytesReceived() float64 {
	return e.MockKibiBytesReceived()
}

func (e *Experiment) KibiBytesSent() float64 {
	return e.MockKibiBytesSent()
}

func (e *Experiment) Name() string {
	return e.MockName()
}

func (e *Experiment) ReportID() string {
	return e.MockReportID()
}

func (e *Experiment) MeasureWithContext(
	ctx context.Context, target model.ExperimentTarget) (measurement *model.Measurement, err error) {
	return e.MockMeasureWithContext(ctx, target)
}

func (e *Experiment) SaveMeasurement(measurement *model.Measurement, filePath string) error {
	return e.MockSaveMeasurement(measurement, filePath)
}

func (e *Experiment) SubmitAndUpdateMeasurementContext(
	ctx context.Context, measurement *model.Measurement) (string, error) {
	return e.MockSubmitAndUpdateMeasurementContext(ctx, measurement)
}

func (e *Experiment) OpenReportContext(ctx context.Context) error {
	return e.MockOpenReportContext(ctx)
}
