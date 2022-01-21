package responsible_person

import (
    "github.com/curatorc/cngf/app"
    "github.com/curatorc/cngf/database"
    "github.com/curatorc/cngf/paginator"

    "github.com/gin-gonic/gin"
)

func Get(idstr string) (responsiblePerson ResponsiblePerson) {
    database.DB.Where("id", idstr).First(&responsiblePerson)
    return
}

func GetBy(field, value string) (responsiblePerson ResponsiblePerson) {
    database.DB.Where("? = ?", field, value).First(&responsiblePerson)
    return
}

func All() (responsiblePeople []ResponsiblePerson) {
    database.DB.Find(&responsiblePeople)
    return
}

func IsExist(field, value string) bool {
    var count int64
    database.DB.Model(ResponsiblePerson{}).Where(" = ?", field, value).Count(&count)
    return count > 0
}

func Paginate(c *gin.Context, perPage int) (responsiblePeople []ResponsiblePerson, paging paginator.Paging) {
    paging = paginator.Paginate(
        c,
        database.DB.Model(ResponsiblePerson{}),
        &responsiblePeople,
        app.V1URL(database.TableName(&ResponsiblePerson{})),
        perPage,
    )
    return
}