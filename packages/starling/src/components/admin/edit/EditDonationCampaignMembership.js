import React from 'react'
import { useQuery } from 'react-apollo-hooks'
import gql from 'graphql-tag'
import PropTypes from 'prop-types'
import Content from '@openmob/bluebird/src/components/layout/Content'
import Card from '@openmob/bluebird/src/components/cards/Card'
import Form from '@openmob/bluebird/src/components/forms/Form'
import Widget from '@openmob/bluebird/src/components/forms/Widget'
import Label from '@openmob/bluebird/src/components/forms/Label'
import Input from '@openmob/bluebird/src/components/forms/Input'
import Button from '@openmob/bluebird/src/components/buttons/Button'

const isObject = a => !!a && a.constructor === Object
const getValue = obj =>
  Object.entries(obj).reduce(entry =>
    entry[0] === 'seconds' || entry[0] === 'ID' ? entry[1] : null
  )
const parseObject = obj => (isObject(obj) ? getValue(obj) : obj)

const GET_DONATIONCAMPAIGNMEMBERSHIP = gql`
  query getDonationCampaignMembershipById($id: ID!) {
    getDonationCampaignMembership(ID: $id) {
      ID
      CreatedAt {
        seconds
      }
      UpdatedAt {
        seconds
      }
      Cause {
        ID
      }
      DonationCampaign {
        ID
      }
    }
  }
`

function EditDonationCampaignMembership({ id }) {
  const {
    data: { getDonationCampaignMembership: item = {} },
    error,
    loading,
  } = useQuery(GET_DONATIONCAMPAIGNMEMBERSHIP, {
    variables: { id },
  })

  if (loading) {
    return <div>Loading...</div>
  }

  if (error) {
    return <div>Error! {error.message}</div>
  }

  return (
    <Content>
      <Card>
        <Form>
          <h1>Edit {item.ID}</h1>
          <Widget>
            <Label>ID</Label>
            <Input value={parseObject(item.ID)} />
          </Widget>
          <Widget>
            <Label>CreatedAt</Label>
            <Input value={parseObject(item.CreatedAt)} />
          </Widget>
          <Widget>
            <Label>UpdatedAt</Label>
            <Input value={parseObject(item.UpdatedAt)} />
          </Widget>
          <Widget>
            <Label>Cause</Label>
            <Input value={parseObject(item.Cause)} />
          </Widget>
          <Widget>
            <Label>DonationCampaign</Label>
            <Input value={parseObject(item.DonationCampaign)} />
          </Widget>

          <Button label="Edit" block variant="primary" />
        </Form>
      </Card>
    </Content>
  )
}

EditDonationCampaignMembership.propTypes = {
  id: PropTypes.string,
}

export default EditDonationCampaignMembership
