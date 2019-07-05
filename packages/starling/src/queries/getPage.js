import gql from 'graphql-tag';

const GET_PAGE = gql`
  query HomePageQuery($id: ID!) {
    getHomePage(ID: $id) {
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
`;

export { GET_PAGE };
