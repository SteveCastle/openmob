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

const MILLISECONDS = 1000
const isObject = a => !!a && a.constructor === Object
const getValue = obj =>
  Object.entries(obj).reduce((acc, entry) => {
    debugger
    if (entry[0] === 'seconds') {
      return new Date(entry[1] * MILLISECONDS)
    }
    if (entry[0] === 'ID') {
      return entry[1]
    }
    return acc
  }, '')
const parseObject = obj => (isObject(obj) ? getValue(obj) : obj)

const GET_CAUSE = gql`
  query getCauseById($id: ID!) {
    getCause(ID: $id) {
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
      Slug
      Summary
      HomePage {
        ID
      }
      Photo {
        ID
      }
    }
  }
`

function EditCause({ id }) {
  const {
    data: { getCause: item = {} },
    error,
    loading,
  } = useQuery(GET_CAUSE, {
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
            <Input value={parseObject(item.ID)} disabled />
          </Widget>
          <Widget>
            <Label>CreatedAt</Label>
            <Input value={parseObject(item.CreatedAt)} disabled />
          </Widget>
          <Widget>
            <Label>UpdatedAt</Label>
            <Input value={parseObject(item.UpdatedAt)} disabled />
          </Widget>
          <Widget>
            <Label>Title</Label>
            <Input value={parseObject(item.Title)} />
          </Widget>
          <Widget>
            <Label>Slug</Label>
            <Input value={parseObject(item.Slug)} />
          </Widget>
          <Widget>
            <Label>Summary</Label>
            <Input value={parseObject(item.Summary)} />
          </Widget>
          <Widget>
            <Label>HomePage</Label>
            <Input value={parseObject(item.HomePage)} />
          </Widget>
          <Widget>
            <Label>Photo</Label>
            <Input value={parseObject(item.Photo)} />
          </Widget>

          <Button label="Edit" block variant="primary" />
        </Form>
      </Card>
    </Content>
  )
}

EditCause.propTypes = {
  id: PropTypes.string,
}

export default EditCause
