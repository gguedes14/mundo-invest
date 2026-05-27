package dto

type ClientInput struct {
	ClienteNome     string  `json:"cliente_nome"`
	ClienteEmail    string  `json:"cliente_email"`
	TipoSolicitacao string  `json:"tipo_solicitacao"`
	ValorPatrimonio float64 `json:"valor_patrimonio"`
}

type ClientResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}
