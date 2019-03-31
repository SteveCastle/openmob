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

const GET_FIELDTYPE = gql`
  query getFieldTypeById($id: ID!) {
    getFieldType(ID: $id) {
      ID
      CreatedAt {
        seconds
        nanos
      }
      UpdatedAt {
        seconds
        nanos
      }
      Title
      DataType
      PropName
      StringValueDefault
      IntValueDefault
      FloatValueDefault
      BooleanValueDefault
      DateTimeValueDefault {
        seconds
        nanos
      }
      ComponentType {
        ID
      }
    }
  }
`

function EditFieldType({ id }) {
  const {
    data: { getFieldType: item = {} },
    error,
    loading,
  } = useQuery(GET_FIELDTYPE, {
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
            <Label>Title</Label>
            <Input value={parseObject(item.Title)} />
          </Widget>
          <Widget>
            <Label>DataType</Label>
            <Input value={parseObject(item.DataType)} />
          </Widget>
          <Widget>
            <Label>PropName</Label>
            <Input value={parseObject(item.PropName)} />
          </Widget>
          <Widget>
            <Label>StringValueDefault</Label>
            <Input value={parseObject(item.StringValueDefault)} />
          </Widget>
          <Widget>
            <Label>IntValueDefault</Label>
            <Input value={parseObject(item.IntValueDefault)} />
          </Widget>
          <Widget>
            <Label>FloatValueDefault</Label>
            <Input value={parseObject(item.FloatValueDefault)} />
          </Widget>
          <Widget>
            <Label>BooleanValueDefault</Label>
            <Input value={parseObject(item.BooleanValueDefault)} />
          </Widget>
          <Widget>
            <Label>DateTimeValueDefault</Label>
            <Input value={parseObject(item.DateTimeValueDefault)} />
          </Widget>
          <Widget>
            <Label>ComponentType</Label>
            <Input value={parseObject(item.ComponentType)} />
          </Widget>

          <Button label="Edit" block variant="primary" />
        </Form>
      </Card>
    </Content>
  )
}

EditFieldType.propTypes = {
  id: PropTypes.string,
}

export default EditFieldType
