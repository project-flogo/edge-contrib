package mqtt

type Settings struct {
	Broker    string `md:"broker"`
	Id        string `md:"id"`
	User      string `md:"user"`
	Password  string `md:"password"`
	Store     string `md:"store"`
	Cleansess bool   `md:"cleansess"`
	Topic     string `md:"topic"`
	Qos       int    `md:"qos"`
}

type Input struct {
	Message interface{} `md:"message"`
}

type Output struct {
	Data interface{} `md:"data"`
}

func (i *Input) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"message": i.Message,
	}
}

func (i *Input) FromMap(values map[string]interface{}) error {
	i.Message, _ = values["message"]
	return nil
}

func (o *Output) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"data": o.Data,
	}
}

func (o *Output) FromMap(values map[string]interface{}) error {

	o.Data = values["data"]
	return nil
}
