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

const GET_EVENTATTENDEE = gql`
  query getEventAttendeeById($id: ID!) {
    getEventAttendee(ID: $id) {
      ID
      CreatedAt {
        seconds
      }
      UpdatedAt {
        seconds
      }
      LiveEvent {
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

function EditEventAttendee({ id }) {
  const {
    data: { getEventAttendee: item = {} },
    error,
    loading,
  } = useQuery(GET_EVENTATTENDEE, {
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
            <Label>LiveEvent</Label>
            <Input value={parseObject(item.LiveEvent)} />
          </Widget>
          <Widget>
            <Label>Contact</Label>
            <Input value={parseObject(item.Contact)} />
          </Widget>
          <Widget>
            <Label>Cause</Label>
            <Input value={parseObject(item.Cause)} />
          </Widget>

          <Button label="Edit" block variant="primary" />
        </Form>
      </Card>
    </Content>
  )
}

EditEventAttendee.propTypes = {
  id: PropTypes.string,
}

export default EditEventAttendee
