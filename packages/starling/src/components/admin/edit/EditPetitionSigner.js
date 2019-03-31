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

const GET_PETITIONSIGNER = gql`
  query getPetitionSignerById($id: ID!) {
    getPetitionSigner(ID: $id) {
      ID
      CreatedAt {
        seconds
        nanos
      }
      UpdatedAt {
        seconds
        nanos
      }
      Petition {
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

function EditPetitionSigner({ id }) {
  const {
    data: { getPetitionSigner: item = {} },
    error,
    loading,
  } = useQuery(GET_PETITIONSIGNER, {
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
            <Label>Petition</Label>
            <Input value={parseObject(item.Petition)} />
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

EditPetitionSigner.propTypes = {
  id: PropTypes.string,
}

export default EditPetitionSigner
