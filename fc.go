package pochtatrack

import (
	"encoding/xml"

	"github.com/widgetii/gowsdl/soap"
	"github.com/widgetii/pochtatrack/fc"
)

type Ticket string

type FCService struct {
	fc       fc.FederalClient
	login    string
	password string
}

type SOAPEnvelopeCustom struct {
	XMLName xml.Name      `xml:"SOAP-ENV:Envelope"`
	XmlSOAP string        `xml:"xmlns:SOAP-ENV,attr"`
	XmlNS0  string        `xml:"xmlns:ns0,attr"`
	Headers []interface{} `xml:"http://schemas.xmlsoap.org/soap/envelope/ Header"`
	Body    SOAPBodyCustom
}

func (s *SOAPEnvelopeCustom) SetHeaders(headers []interface{}) {
	s.Headers = headers
}

func (s *SOAPEnvelopeCustom) SetContent(content interface{}) {
	s.Body = SOAPBodyCustom{Content: content}
}

func (s SOAPEnvelopeCustom) Fault() *soap.SOAPFault {
	return s.Body.Fault
}

type SOAPBodyCustom struct {
	XMLName xml.Name `xml:"SOAP-ENV:Body"`

	Fault   *soap.SOAPFault `xml:",omitempty"`
	Content interface{}     `xml:",omitempty"`
}

func (s SOAPBodyCustom) content() interface{} {
	return s.Content
}

func (s *SOAPBodyCustom) setContent(c interface{}) {
	s.Content = c
}

func (s SOAPBodyCustom) fault() *soap.SOAPFault {
	return s.Fault
}

func (s *SOAPBodyCustom) setFault(f *soap.SOAPFault) {
	s.Fault = f
}

func customSOAPRequester() soap.SOAPEnvelope {
	return &SOAPEnvelopeCustom{XmlSOAP: "http://schemas.xmlsoap.org/soap/envelope/",
		XmlNS0: "http://fclient.russianpost.org/postserver"}
}

func NewFC(login string, password string) *FCService {
	c := soap.NewClient("https://tracking.russianpost.ru/fc",
		soap.WithCustomRequester(customSOAPRequester),
		soap.WithSOAPVersion(soap.SOAPVersion(soap.SOAP11)))

	return &FCService{
		fc:       fc.NewFederalClient(c),
		login:    login,
		password: password}
}

func (f *FCService) GetTicket([]string) (Ticket, error) {
	items := []*fc.Item{
		&fc.Item{Barcode: "19003132224427"},
		&fc.Item{Barcode: "12727630039983"},
	}
	res, err := f.fc.GetTicket(&fc.TicketRequest{
		Request: &fc.File{
			Item: items,
		},
		Login:    f.login,
		Password: f.password,
		Language: "RUS",
	})
	if err != nil {
		return Ticket(""), err
	}
	return Ticket(res.Value), nil
}

func (f *FCService) GetResponse(ticket Ticket) (*fc.AnswerByTicketResponse, error) {
	return f.fc.GetResponseByTicket(&fc.AnswerByTicketRequest{
		Login:    f.login,
		Password: f.password,
		Ticket:   string(ticket),
	})
}
