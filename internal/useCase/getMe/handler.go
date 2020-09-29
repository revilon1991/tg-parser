package getMe

import "github.com/revilon1991/tg-parser/internal/client"

func Handle(clientStorage *client.Storage) Me {
    var me Me

    res, _ := clientStorage.SendAndCatch(client.Request{
        "@type": "getMe",
    })

    me.Data = res

    return me
}
