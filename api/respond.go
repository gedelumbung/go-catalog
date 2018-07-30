package api

type (
	Respond struct {
		Data interface{} `json:"data"`
		Meta interface{} `json:"meta,omitempty"`
	}

	RespondError struct {
		Scope  string      `json:"scope"`
		Reason string      `json:"reason"`
		Detail interface{} `json:"detail,omitempty"`
	}
)

func OKRespond(data interface{}) Respond {
	return Respond{Data: data}
}

func OKRespondWithMeta(data interface{}, meta interface{}) Respond {
	return Respond{Data: data, Meta: meta}
}

func ErrRespond(scope, message string, err interface{}) Respond {
	return Respond{
		Data: RespondError{
			Scope:  scope,
			Reason: message,
			Detail: err,
		},
	}
}

func ErrRespondString(scope, message, err string) Respond {
	return Respond{
		Data: RespondError{
			Scope:  scope,
			Reason: message,
			Detail: err,
		},
	}
}
