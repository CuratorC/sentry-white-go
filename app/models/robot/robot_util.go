package robot

import (
    "github.com/curatorc/cngf/app"
    "github.com/curatorc/cngf/database"
    "github.com/curatorc/cngf/paginator"

    "github.com/gin-gonic/gin"
)

func Get(idstr string) (robot Robot) {
    database.DB.Where("id", idstr).First(&robot)
    return
}

func GetBy(field, value string) (robot Robot) {
    database.DB.Where("? = ?", field, value).First(&robot)
    return
}

func All() (robots []Robot) {
    database.DB.Find(&robots)
    return
}

func IsExist(field, value string) bool {
    var count int64
    database.DB.Model(Robot{}).Where(" = ?", field, value).Count(&count)
    return count > 0
}

func Paginate(c *gin.Context, perPage int) (robots []Robot, paging paginator.Paging) {
    paging = paginator.Paginate(
        c,
        database.DB.Model(Robot{}),
        &robots,
        app.V1URL(database.TableName(&Robot{})),
        perPage,
    )
    return
}