package enum

import (
    "fmt"
    "slices"
    "strings"
)

// This file has been created automatically by `go-enum-generate`
// DO NOT MODIFY NOR EDIT THIS FILE DIRECTLY.
// To modify this enum, edit the enums.json or enums.yaml definition file
// To know more about `go-enum-generate`, see go to `https://github.com/debarbarinantoine/go-enum-generate`
// Generated at: 2025-08-12 11:46:33

type FileType uint

const (
    atom FileType = iota
    molecule
    organism
    template
    model
    provider
    repository
    service
    util
    page
)

var fileTypeKeys = make(map[FileType]struct{}, 10)
var fileTypeValues = make(map[string]FileType, 10)
var fileTypeKeysArray = make([]FileType, 10)
var fileTypeValuesArray = make([]string, 10)

func init() {
    fileTypeKeys[atom] = struct{}{}
    fileTypeKeysArray[0] = atom
    fileTypeValues["components/atoms"] = atom
    fileTypeValuesArray[0] = "components/atoms"

    fileTypeKeys[molecule] = struct{}{}
    fileTypeKeysArray[1] = molecule
    fileTypeValues["components/molecules"] = molecule
    fileTypeValuesArray[1] = "components/molecules"

    fileTypeKeys[organism] = struct{}{}
    fileTypeKeysArray[2] = organism
    fileTypeValues["components/organisms"] = organism
    fileTypeValuesArray[2] = "components/organisms"

    fileTypeKeys[template] = struct{}{}
    fileTypeKeysArray[3] = template
    fileTypeValues["components/templates"] = template
    fileTypeValuesArray[3] = "components/templates"

    fileTypeKeys[model] = struct{}{}
    fileTypeKeysArray[4] = model
    fileTypeValues["lib/models"] = model
    fileTypeValuesArray[4] = "lib/models"

    fileTypeKeys[provider] = struct{}{}
    fileTypeKeysArray[5] = provider
    fileTypeValues["lib/providers"] = provider
    fileTypeValuesArray[5] = "lib/providers"

    fileTypeKeys[repository] = struct{}{}
    fileTypeKeysArray[6] = repository
    fileTypeValues["lib/repositories"] = repository
    fileTypeValuesArray[6] = "lib/repositories"

    fileTypeKeys[service] = struct{}{}
    fileTypeKeysArray[7] = service
    fileTypeValues["lib/services"] = service
    fileTypeValuesArray[7] = "lib/services"

    fileTypeKeys[util] = struct{}{}
    fileTypeKeysArray[8] = util
    fileTypeValues["lib/utils"] = util
    fileTypeValuesArray[8] = "lib/utils"

    fileTypeKeys[page] = struct{}{}
    fileTypeKeysArray[9] = page
    fileTypeValues["page"] = page
    fileTypeValuesArray[9] = "page"
}

func (e FileType) String() string {
    switch e {
        case atom:
            return "components/atoms"
        case molecule:
            return "components/molecules"
        case organism:
            return "components/organisms"
        case template:
            return "components/templates"
        case model:
            return "lib/models"
        case provider:
            return "lib/providers"
        case repository:
            return "lib/repositories"
        case service:
            return "lib/services"
        case util:
            return "lib/utils"
        case page:
            return "page"
        default:
            return fmt.Sprintf("Unknown FileType (%d)", e.Value())
    }
}

func (e *FileType) Parse(str string) error {

    str = strings.TrimSpace(str)

    if val, ok := fileTypeValues[str]; ok {
        *e = val
        return nil
    }
    return fmt.Errorf("invalid FileType: %s", str)
}

func (e FileType) Value() uint {
    return uint(e)
}

func (e FileType) MarshalText() ([]byte, error) {
    return []byte(e.String()), nil
}

func (e *FileType) UnmarshalText(text []byte) error {
    return e.Parse(string(text))
}

func (e FileType) IsValid() bool {
    if _, ok := fileTypeKeys[e]; !ok {
        return false
    }
    return true
}

type fileTypes struct {
    Atom FileType
    Molecule FileType
    Organism FileType
    Template FileType
    Model FileType
    Provider FileType
    Repository FileType
    Service FileType
    Util FileType
    Page FileType
}

var FileTypes = fileTypes{
    Atom: atom,
    Molecule: molecule,
    Organism: organism,
    Template: template,
    Model: model,
    Provider: provider,
    Repository: repository,
    Service: service,
    Util: util,
    Page: page,
}

func (e fileTypes) Values() []FileType {
    return slices.Clone(fileTypeKeysArray)
}

func (e fileTypes) Args() []string {
    return slices.Clone(fileTypeValuesArray)
}

func (e fileTypes) Description() string {
    var strBuilder strings.Builder
    strBuilder.WriteString("\tAvailable FileTypes:\n")
    for _, enumVal := range e.Values() {
        strBuilder.WriteString(fmt.Sprintf("=> %d -> %s\n", enumVal.Value(), enumVal.String()))
    }
    return strBuilder.String()
}

func (e fileTypes) Cast(value uint) (FileType, error) {
    if _, ok := fileTypeKeys[FileType(value)]; !ok {
        return 0, fmt.Errorf("invalid cast FileType: %d", value)
    }
    return FileType(value), nil
}
