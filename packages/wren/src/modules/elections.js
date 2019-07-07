// This module maps a dataPath field in a layout
// to a set of gRPC requests to load data required for the dataPath.
// This will probably be moved into gRPC service.

async function getFieldValue(client, causeId) {
  const memberData = await client
    .ListElectionMembership()
    .sendMessage({ api: 'v1', filters: [{ Cause: causeId }], limit: 10 });
  const memberItems = memberData.items;

  const fetchItems = async item => {
    const result = await client
      .GetElection()
      .sendMessage({ api: 'v1', ID: item.Election });
    return result.item;
  };
  const getData = async () => {
    return await Promise.all(memberItems.map(item => fetchItems(item)));
  };
  const data = await getData();
  return JSON.stringify(data);
}

module.exports = {
  getFieldValue
};
