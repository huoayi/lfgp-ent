package enum

type CreationType string

const (
	Txt2voice   CreationType = "txt2voice"
	Txt2img     CreationType = "txt2img"
	Voice2voice CreationType = "voice2voice"
)

func (obj CreationType) Values() []string {
	return []string{
		string(Txt2voice),
		string(Txt2img),
		string(Voice2voice),
	}
}
func (obj CreationType) Ptr() *CreationType {
	if obj != "" {
		return &obj
	} else {
		return nil
	}
}
