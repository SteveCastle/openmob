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

const GET_FIELD = gql`
  query getFieldById($id: ID!) {
    getField(ID: $id) {
      ID
      CreatedAt {
        seconds
      }
      UpdatedAt {
        seconds
      }
      FieldType {
        ID
      }
      StringValue
      IntValue
      FloatValue
      BooleanValue
      DateTimeValue {
        seconds
      }
      Component {
        ID
      }
    }
  }
`

function EditField({ id }) {
  const {
    data: { getField: item = {} },
    error,
    loading,
  } = useQuery(GET_FIELD, {
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
            <Label>FieldType</Label>
            <Input value={parseObject(item.FieldType)} />
          </Widget>
          <Widget>
            <Label>StringValue</Label>
            <Input value={parseObject(item.StringValue)} />
          </Widget>
          <Widget>
            <Label>IntValue</Label>
            <Input value={parseObject(item.IntValue)} />
          </Widget>
          <Widget>
            <Label>FloatValue</Label>
            <Input value={parseObject(item.FloatValue)} />
          </Widget>
          <Widget>
            <Label>BooleanValue</Label>
            <Input value={parseObject(item.BooleanValue)} />
          </Widget>
          <Widget>
            <Label>DateTimeValue</Label>
            <Input value={parseObject(item.DateTimeValue)} />
          </Widget>
          <Widget>
            <Label>Component</Label>
            <Input value={parseObject(item.Component)} />
          </Widget>

          <Button label="Edit" block variant="primary" />
        </Form>
      </Card>
    </Content>
  )
}

EditField.propTypes = {
  id: PropTypes.string,
}

export default EditField
