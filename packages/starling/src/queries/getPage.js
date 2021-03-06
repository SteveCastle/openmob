import gql from 'graphql-tag';

const GET_PAGE = gql`
  query HomePageQuery($id: ID!) {
    getCause(ID: $id) {
      HomePage {
        ID
        Title
        Layout {
          ID
          LayoutRows {
            ID
            Container
            Weight
            LayoutColumns {
              ID
              Width
              Weight
              Components {
                ID
                Weight
                Fields {
                  ID
                  FieldType {
                    ID
                    Title
                    DataType
                    PropName
                  }
                  StringValue
                  IntValue
                  FloatValue
                  BooleanValue
                  DateTimeValue {
                    seconds
                  }
                  DataPathValue
                  DataPath
                }
                ComponentType {
                  ID
                  ComponentTypeFieldss {
                    ID
                    FieldType {
                      ID
                      Title
                    }
                  }
                  ComponentImplementations {
                    ID
                    Title
                  }
                  ComponentImplementation {
                    ID
                  }
                }
                ComponentImplementation {
                  ID
                  Title
                  Path
                }
              }
            }
          }
          LayoutType {
            ID
          }
        }
      }
    }
  }
`;

export { GET_PAGE };
