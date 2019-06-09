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
              ComponentImplementation {
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

export default GET_PAGE;
