package pipefy

const CreateCardMutation = `
mutation {
	createCard(
		input: {
			pipe_id: 123456
			title: "%s"
			fields_attributes: [
				{
					field_id: "cliente_nome"
					field_value: "%s"
				},
				{
					field_id: "cliente_email"
					field_value: "%s"
				},
				{
					field_id: "valor_patrimonio"
					field_value: "%.2f"
				}
			]
		}
	) {
		card {
			id
			title
		}
	}
}
`
