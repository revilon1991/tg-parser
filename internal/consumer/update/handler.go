package update

import (
    "github.com/revilon1991/tg-parser/internal/client"
    "github.com/revilon1991/tg-parser/internal/useCase/authorization"
    "github.com/revilon1991/tg-parser/internal/useCase/channelCatch"
    "log"
)

func Handle(clientStorage *client.Storage) {
    go func() {
        for update := range clientStorage.Updates {
            //Print all updates
            //spew.Dump(update)

            if update.UpdateData["@type"].(string) == "updateAuthorizationState" {
                if authorizationState, ok := update.UpdateData["authorization_state"].(map[string]interface{})["@type"].(string); ok {
                    clientStorage.AuthorizationState = authorizationState

                    _, err := authorization.Handle(clientStorage, authorizationState)

                    if err != nil {
                        log.Println("updateAuthorizationState error: " + err.Error())
                    }
                }
            }

            if update.UpdateData["@type"].(string) == "updateSupergroup" {
                channelCatch.Handle(clientStorage, update)
            }
        }
    }()
}
