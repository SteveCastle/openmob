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
  console.log(data);
  return JSON.stringify(data);
}

module.exports = {
  getFieldValue,
};
