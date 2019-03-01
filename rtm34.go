package pochtatrack

import (
	"github.com/widgetii/gowsdl/soap"
	"github.com/widgetii/pochtatrack/rtm34"
)

type RTM34Service struct {
	opHistory rtm34.OperationHistory12
	creds     *rtm34.AuthorizationHeader
}

func NewRTM34(login string, password string) *RTM34Service {
	c := soap.NewClient("http://tracking.russianpost.ru/rtm34",
		soap.WithSOAPVersion(soap.SOAPVersion(soap.SOAP12)))
	return &RTM34Service{
		opHistory: rtm34.NewOperationHistory12(c),
		creds: &rtm34.AuthorizationHeader{
			Login:    login,
			Password: password,
		}}
}

func (r *RTM34Service) GetOperationHistory(barcode string) (*rtm34.GetOperationHistoryResponse, error) {
	return r.opHistory.GetOperationHistory(
		&rtm34.GetOperationHistory{
			AuthorizationHeader: r.creds,
			OperationHistoryRequest: &rtm34.OperationHistoryRequest{
				Barcode: barcode,
			},
		})
}
