package utils

type MessagePrototype struct {
	APIVersion  string       `json:"apiVersion"`
	DataSuccess *DataObject  `json:"dataSuccess,omitempty"`
	DataError   *ErrorObject `json:"dataError,omitempty"`
}

type DataObject struct {
	Kind        *string     `json:"kind,omitempty"`
	Title       string      `json:"title,omitempty"`
	Description string      `json:"description,omitempty"`
	Item        interface{} `json:"item,omitempty"`
	Items       interface{} `json:"items,omitempty"`
}

type ErrorObject struct {
	Title        string `json:"title,omitempty"`
	ErrorMessage string `json:"description,omitempty"`
}

func SuccessMessage(data DataObject) MessagePrototype {
	return MessagePrototype{APIVersion: "v1", DataSuccess: &data}
}

func ErrorMessage(data ErrorObject) MessagePrototype {
	return MessagePrototype{APIVersion: "v1", DataError: &data}
}
