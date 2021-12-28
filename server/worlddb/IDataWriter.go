package worlddb

type IDataWriter interface {
	SaveModifyToDB() bool
}
