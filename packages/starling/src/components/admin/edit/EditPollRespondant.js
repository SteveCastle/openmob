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

const GET_POLLRESPONDANT = gql`
  query getPollRespondantById($id: ID!) {
    getPollRespondant(ID: $id) {
      ID
      CreatedAt {
        seconds
      }
      UpdatedAt {
        seconds
      }
      Poll {
        ID
      }
      Contact {
        ID
      }
      Cause {
        ID
      }
    }
  }
`

function EditPollRespondant({ id }) {
  const {
    data: { getPollRespondant: item = {} },
    error,
    loading,
  } = useQuery(GET_POLLRESPONDANT, {
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
            <Label>Poll</Label>
            <Input placeholder={parseObject(item.Poll)} />
          </Widget>
          <Widget>
            <Label>Contact</Label>
            <Input placeholder={parseObject(item.Contact)} />
          </Widget>
          <Widget>
            <Label>Cause</Label>
            <Input placeholder={parseObject(item.Cause)} />
          </Widget>

          <Button label="Edit" block />
        </Form>
      </Card>
    </Content>
  )
}

EditPollRespondant.propTypes = {
  id: PropTypes.string,
}

export default EditPollRespondant
