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

const GET_ELECTIONMEMBERSHIP = gql`
  query getElectionMembershipById($id: ID!) {
    getElectionMembership(ID: $id) {
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
      Election {
        ID
      }
    }
  }
`

function EditElectionMembership({ id }) {
  const {
    data: { getElectionMembership: item = {} },
    error,
    loading,
  } = useQuery(GET_ELECTIONMEMBERSHIP, {
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
            <Input placeholder={parseObject(item.ID)} />
          </Widget>
          <Widget>
            <Label>CreatedAt</Label>
            <Input placeholder={parseObject(item.CreatedAt)} />
          </Widget>
          <Widget>
            <Label>UpdatedAt</Label>
            <Input placeholder={parseObject(item.UpdatedAt)} />
          </Widget>
          <Widget>
            <Label>Cause</Label>
            <Input placeholder={parseObject(item.Cause)} />
          </Widget>
          <Widget>
            <Label>Election</Label>
            <Input placeholder={parseObject(item.Election)} />
          </Widget>

          <Button label="Edit" />
        </Form>
      </Card>
    </Content>
  )
}

EditElectionMembership.propTypes = {
  id: PropTypes.string,
}

export default EditElectionMembership
