// TODO: Move this to shrike when behavior is defined.
// We would prefer not to have this logic in the graphQL server
// but it's a good place to experiment with features quickly.

const { mergeResolvers } = require('merge-graphql-schemas');
const generatedResolvers = require('./generated/resolvers.js');
const modules = require('./modules');

// Resolvers object is passed grpc client to use for data fetching.
const resolvers = client => ({
  Field: {
    // If there is a DataPath value in the field.
    // Fetch the Membership entries from the CMS and return as json on DataPathValue field.
    DataPathValue: async ({ DataPath }, _, ctx) => {
      console.log('resolving datapathvalue: ', ctx.cause, DataPath);
      if (DataPath && ctx.cause) {
        const data = await modules[DataPath].getFieldValue(client, ctx.cause);
        return data;
      }
      return null;
    }
  },
  Query: {
    getCause: async (_, { ID }, ctx) => {
      ctx.cause = ID;
      const data = await client.GetCause().sendMessage({ api: 'v1', ID });
      return data.item;
    }
  },
  Mutation: {
    // Create a cause and all of its mandatory starting relations.
    newCause: async (
      _,
      {
        Title,
        Slug,
        Summary,
        FeaturedImage = 'https://punknaturalism.com/static/84e6008a85bd56160b3e6d31f23d428e/a5547/N33fe4iyc6.jpg'
      }
    ) => {
      const { ID: layoutID } = await client.CreateLayout().sendMessage({
        api: 'v1',
        item: { LayoutType: 'adb5a57a-b0ab-4022-8d14-bde3efbe5ad9' }
      });
      const { ID: homePageID } = await client.CreateHomePage().sendMessage({
        api: 'v1',
        item: { Title: 'Home Page', Layout: layoutID }
      });
      const { ID: photoID } = await client.CreatePhoto().sendMessage({
        api: 'v1',
        item: {
          URI: FeaturedImage,
          Width: 400,
          Height: 400
        }
      });
      const { ID: causeID } = await client.CreateCause().sendMessage({
        api: 'v1',

        item: {
          Title,
          Slug,
          Summary,
          HomePage: homePageID,
          Photo: photoID
        },
        buildStatic: true
      });
      return causeID;
    },
    // Create a new component of a certain type.
    // with the fields defined on its type.
    newComponent: async (
      _,
      { LayoutColumn, ComponentType, ComponentImplementation }
    ) => {
      const { ID: componentID } = await client.CreateComponent().sendMessage({
        api: 'v1',

        item: {
          LayoutColumn,
          ComponentType,
          ComponentImplementation
        },
        buildStatic: false
      });
      const { items } = await client.ListComponentTypeFields().sendMessage({
        api: 'v1',
        filters: [{ ComponentType: ComponentType }],
        limit: 10
      });
      if (!items) {
        return componentID;
      }
      const createField = async item => {
        const { item: fieldType } = await client
          .GetFieldType()
          .sendMessage({ api: 'v1', ID: item.FieldType });
        const field = {
          Component: componentID,
          FieldType: fieldType.ID
        };
        await client
          .CreateField()
          .sendMessage({ api: 'v1', item: { ...field }, buildStatic: false });
      };
      const createFields = async () => {
        return await Promise.all(items.map(item => createField(item)));
      };

      await createFields();
      return componentID;
    }
  }
});

// Merge the generated resolvers with our custom resolvers and export, passing gRPC client along.
const resolversList = client => [generatedResolvers(client), resolvers(client)];
module.exports = client => mergeResolvers(resolversList(client));
