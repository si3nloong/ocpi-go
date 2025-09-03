package ocpi230

import (
	"encoding/json"
	"net/http"

	"github.com/si3nloong/ocpi-go/ocpi"
	ocpihttp "github.com/si3nloong/ocpi-go/ocpi/http"
)

func (s *Server) GetOcpiPtpPaymentTerminals(w http.ResponseWriter, r *http.Request) {
	params := GetPaymentTerminalsParams{}
	response, err := s.paymentsSender.OnGetPaymentTerminals(r.Context(), params)
	if err != nil {
		ocpihttp.Response(w, err)
		return
	}

	ocpihttp.ResponsePagination(w, r, response)
}

func (s *Server) GetOcpiPtpPaymentTerminal(w http.ResponseWriter, r *http.Request) {
	terminalID := r.PathValue("terminal_id")
	terminal, err := s.paymentsSender.OnGetPaymentTerminal(r.Context(), terminalID)
	if err != nil {
		ocpihttp.Response(w, err)
		return
	}

	ocpihttp.Response(w, terminal)
}

func (s *Server) PostOcpiPtpActivatePaymentTerminal(w http.ResponseWriter, r *http.Request) {
	var body json.RawMessage
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		ocpihttp.BadRequest(w, r, err.Error())
		return
	}

	if err := s.paymentsSender.OnPostActivatePaymentTerminal(r.Context(), ocpi.RawMessage[Terminal](body)); err != nil {
		ocpihttp.Response(w, err)
		return
	}

	ocpihttp.Response(w, ocpi.NewEmptyResponse())
}

func (s *Server) PostOcpiPtpDeactivatePaymentTerminal(w http.ResponseWriter, r *http.Request) {
	terminalID := r.PathValue("terminal_id")
	if err := s.paymentsSender.OnPostDeactivatePaymentTerminal(r.Context(), terminalID); err != nil {
		ocpihttp.Response(w, err)
		return
	}

	ocpihttp.Response(w, ocpi.NewEmptyResponse())
}

func (s *Server) PutOcpiPtpPaymentTerminal(w http.ResponseWriter, r *http.Request) {
	var body json.RawMessage
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		ocpihttp.BadRequest(w, r, err.Error())
		return
	}

	if err := s.paymentsSender.OnPutPaymentTerminal(r.Context(), ocpi.RawMessage[Terminal](body)); err != nil {
		ocpihttp.Response(w, err)
		return
	}

	ocpihttp.Response(w, ocpi.NewEmptyResponse())
}

func (s *Server) PatchOcpiPtpPaymentTerminal(w http.ResponseWriter, r *http.Request) {
	var body json.RawMessage
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		ocpihttp.BadRequest(w, r, err.Error())
		return
	}

	if err := s.paymentsSender.OnPatchPaymentTerminal(r.Context(), ocpi.RawMessage[PartialTerminal](body)); err != nil {
		ocpihttp.Response(w, err)
		return
	}

	ocpihttp.Response(w, ocpi.NewEmptyResponse())
}

func (s *Server) GetOcpiPtpFinancialAdviceConfirmations(w http.ResponseWriter, r *http.Request) {
	params := GetPaymentFinancialAdviceConfirmationsParams{}
	response, err := s.paymentsSender.OnGetPaymentFinancialAdviceConfirmations(r.Context(), params)
	if err != nil {
		ocpihttp.Response(w, err)
		return
	}

	ocpihttp.ResponsePagination(w, r, response)
}

func (s *Server) GetOcpiPtpFinancialAdviceConfirmation(w http.ResponseWriter, r *http.Request) {
	financialAdviceConfirmationID := r.PathValue("financial_advice_confirmation_id")
	financialAdviceConfirmation, err := s.paymentsSender.OnGetPaymentFinancialAdviceConfirmation(r.Context(), financialAdviceConfirmationID)
	if err != nil {
		ocpihttp.Response(w, err)
		return
	}

	ocpihttp.Response(w, financialAdviceConfirmation)
}

func (s *Server) GetOcpiPaymentTerminal(w http.ResponseWriter, r *http.Request) {
	terminalID := r.PathValue("terminal_id")
	terminal, err := s.paymentsReceiver.OnGetTerminal(r.Context(), terminalID)
	if err != nil {
		ocpihttp.Response(w, err)
		return
	}

	ocpihttp.Response(w, terminal)
}

func (s *Server) PostOcpiPaymentTerminal(w http.ResponseWriter, r *http.Request) {
	var body json.RawMessage
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		ocpihttp.BadRequest(w, r, err.Error())
		return
	}

	if err := s.paymentsReceiver.OnPostTerminal(r.Context(), ocpi.RawMessage[Terminal](body)); err != nil {
		ocpihttp.Response(w, err)
		return
	}

	ocpihttp.Response(w, ocpi.NewEmptyResponse())
}

func (s *Server) GetOcpiFinancialAdviceConfirmation(w http.ResponseWriter, r *http.Request) {
	financialAdviceConfirmationID := r.PathValue("financial_advice_confirmation_id")
	financialAdviceConfirmation, err := s.paymentsReceiver.OnGetFinancialAdviceConfirmation(r.Context(), financialAdviceConfirmationID)
	if err != nil {
		ocpihttp.Response(w, err)
		return
	}

	ocpihttp.Response(w, financialAdviceConfirmation)
}

func (s *Server) PostOcpiFinancialAdviceConfirmation(w http.ResponseWriter, r *http.Request) {
	var body json.RawMessage
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		ocpihttp.BadRequest(w, r, err.Error())
		return
	}

	if err := s.paymentsReceiver.OnPostFinancialAdviceConfirmation(r.Context(), ocpi.RawMessage[FinancialAdviceConfirmation](body)); err != nil {
		ocpihttp.Response(w, err)
		return
	}

	ocpihttp.Response(w, ocpi.NewEmptyResponse())
}
