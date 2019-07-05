import gql from 'graphql-tag';
const LIST_COMPONENT_TYPES = gql`
  query ComponentTypeList {
    listComponentType(limit: 20) {
      ID
      Title
      ComponentImplementation {
        ID
      }
    }
  }
`;

export { LIST_COMPONENT_TYPES };
