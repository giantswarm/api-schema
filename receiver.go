package apischema

func IsStatusData(responseBody string) (bool, error) {
  return IsStatus(STATUS_CODE_DATA, responseBody)
}

func IsStatusRessourceCreated(responseBody string) (bool, error) {
  return IsStatus(STATUS_CODE_RESSOURCE_CREATED, responseBody)
}

func IsStatusRessourceStarted(responseBody string) (bool, error) {
  return IsStatus(STATUS_CODE_RESSOURCE_STARTED, responseBody)
}

func IsStatusRessourceStopped(responseBody string) (bool, error) {
  return IsStatus(STATUS_CODE_RESSOURCE_STOPPED, responseBody)
}

func IsStatusRessourceUpdated(responseBody string) (bool, error) {
  return IsStatus(STATUS_CODE_RESSOURCE_UPDATED, responseBody)
}

func IsStatusRessourceDeleted(responseBody string) (bool, error) {
  return IsStatus(STATUS_CODE_RESSOURCE_DELETED, responseBody)
}

func IsStatusRessourceAlreadyExists(responseBody string) (bool, error) {
  return IsStatus(STATUS_CODE_RESSOURCE_ALREADY_EXISTS, responseBody)
}
