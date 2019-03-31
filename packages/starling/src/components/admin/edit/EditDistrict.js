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

const GET_DISTRICT = gql`
  query getDistrictById($id: ID!) {
    getDistrict(ID: $id) {
      ID
      CreatedAt {
        seconds
        nanos
      }
      UpdatedAt {
        seconds
        nanos
      }
      Geom
      Title
      DistrictType {
        ID
      }
    }
  }
`

function EditDistrict({ id }) {
  const {
    data: { getDistrict: item = {} },
    error,
    loading,
  } = useQuery(GET_DISTRICT, {
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
            <Label>Geom</Label>
            <Input value={parseObject(item.Geom)} />
          </Widget>
          <Widget>
            <Label>Title</Label>
            <Input value={parseObject(item.Title)} />
          </Widget>
          <Widget>
            <Label>DistrictType</Label>
            <Input value={parseObject(item.DistrictType)} />
          </Widget>

          <Button label="Edit" block variant="primary" />
        </Form>
      </Card>
    </Content>
  )
}

EditDistrict.propTypes = {
  id: PropTypes.string,
}

export default EditDistrict
