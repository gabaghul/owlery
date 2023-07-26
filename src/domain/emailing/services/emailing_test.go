package services_test

import (
	"context"
	"errors"
	"testing"

	"github.com/gabaghul/owlery/src/domain/emailing/models"
	"github.com/gabaghul/owlery/src/domain/emailing/services"
	mock_services "github.com/gabaghul/owlery/src/domain/emailing/services/mocks"
	"github.com/golang/mock/gomock"
	"github.com/onsi/gomega"
	"github.com/rs/zerolog"
)

func Test_DoEmailPooling(t *testing.T) {
	g := gomega.NewWithT(t)
	ctx := context.Background()
	logger := &zerolog.Logger{}
	ctrl := gomock.NewController(t)

	redis := mock_services.NewMockRedisAdapter(ctrl)
	psql := mock_services.NewMockPsqlAdapter(ctrl)
	mailchimp := mock_services.NewMockMailChimpAdapter(ctrl)
	ometria := mock_services.NewMockOmetriaAdapter(ctrl)

	service := services.NewEmailingService(logger, psql, redis, mailchimp, ometria)
	type testCase struct {
		description string
		beforeTest  func(tc *testCase)
		expectError bool
	}

	// I just did one testcase because I lacked time this week, but the whole idea would be mocking all interfaces to return
	// errors and structs so we can achieve the best coverage possible. Hurrah for mockgen!!
	testCases := []testCase{
		{
			description: "should return error when get emailing configs returns error",
			beforeTest: func(tc *testCase) {
				psql.EXPECT().GetAllEmailingConfigs(ctx).Return([]models.EmailingConfig{}, errors.New("mocked error"))
			},
			expectError: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			if tc.beforeTest != nil {
				tc.beforeTest(&tc)
			}

			err := service.DoEmailPooling(ctx)
			if tc.expectError {
				g.Expect(err).To(gomega.HaveOccurred(), tc.description)
			} else {
				g.Expect(err).ToNot(gomega.HaveOccurred(), tc.description)
			}
		})
	}
}
