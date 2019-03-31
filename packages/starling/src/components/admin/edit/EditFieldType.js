import React from 'react'
import { useQuery, useMutation } from 'react-apollo-hooks'
import gql from 'graphql-tag'
import { Formik } from 'formik'
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
    if (entry[0] === 'seconds') {
      return new Date(entry[1] * MILLISECONDS).toString()
    }
    if (entry[0] === 'ID') {
      return entry[1]
    }
    return acc
  }, '')
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
const UPDATE_FIELDTYPE = gql`
  mutation updateFieldType($id: ID!, $fieldType: FieldTypeInput) {
    updateFieldType(ID: $id, fieldType: $fieldType, buildStatic: false)
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

  const updateFieldType = useMutation(UPDATE_FIELDTYPE)

  if (loading) {
    return <div>Loading...</div>
  }

  if (error) {
    return <div>Error! {error.message}</div>
  }

  return (
    <Formik
      initialValues={{
        ID: parseObject(item.ID),
        CreatedAt: parseObject(item.CreatedAt),
        UpdatedAt: parseObject(item.UpdatedAt),
        Title: parseObject(item.Title),
        DataType: parseObject(item.DataType),
        PropName: parseObject(item.PropName),
        StringValueDefault: parseObject(item.StringValueDefault),
        IntValueDefault: parseObject(item.IntValueDefault),
        FloatValueDefault: parseObject(item.FloatValueDefault),
        BooleanValueDefault: parseObject(item.BooleanValueDefault),
        DateTimeValueDefault: parseObject(item.DateTimeValueDefault),
        ComponentType: parseObject(item.ComponentType),
      }}
      onSubmit={(values, { setSubmitting }) =>
        updateFieldType({
          variables: {
            id: item.ID,
            fieldType: {
              ...values,
              ID: undefined,
              CreatedAt: undefined,
              UpdatedAt: undefined,
            },
          },
        })
      }
    >
      {props => {
        const { values, handleChange, handleBlur, handleSubmit } = props
        return (
          <Content>
            <Card>
              <Form>
                <h1>Edit {item.ID}</h1>
                <Widget>
                  <Label>ID</Label>
                  <Input
                    value={values.ID}
                    disabled
                    name="ID"
                    onChange={handleChange}
                    onBlur={handleBlur}
                  />
                </Widget>
                <Widget>
                  <Label>CreatedAt</Label>
                  <Input
                    value={values.CreatedAt}
                    disabled
                    name="CreatedAt"
                    onChange={handleChange}
                    onBlur={handleBlur}
                  />
                </Widget>
                <Widget>
                  <Label>UpdatedAt</Label>
                  <Input
                    value={values.UpdatedAt}
                    disabled
                    name="UpdatedAt"
                    onChange={handleChange}
                    onBlur={handleBlur}
                  />
                </Widget>
                <Widget>
                  <Label>Title</Label>
                  <Input
                    value={values.Title}
                    name="Title"
                    onChange={handleChange}
                    onBlur={handleBlur}
                  />
                </Widget>
                <Widget>
                  <Label>DataType</Label>
                  <Input
                    value={values.DataType}
                    name="DataType"
                    onChange={handleChange}
                    onBlur={handleBlur}
                  />
                </Widget>
                <Widget>
                  <Label>PropName</Label>
                  <Input
                    value={values.PropName}
                    name="PropName"
                    onChange={handleChange}
                    onBlur={handleBlur}
                  />
                </Widget>
                <Widget>
                  <Label>StringValueDefault</Label>
                  <Input
                    value={values.StringValueDefault}
                    name="StringValueDefault"
                    onChange={handleChange}
                    onBlur={handleBlur}
                  />
                </Widget>
                <Widget>
                  <Label>IntValueDefault</Label>
                  <Input
                    value={values.IntValueDefault}
                    name="IntValueDefault"
                    onChange={handleChange}
                    onBlur={handleBlur}
                  />
                </Widget>
                <Widget>
                  <Label>FloatValueDefault</Label>
                  <Input
                    value={values.FloatValueDefault}
                    name="FloatValueDefault"
                    onChange={handleChange}
                    onBlur={handleBlur}
                  />
                </Widget>
                <Widget>
                  <Label>BooleanValueDefault</Label>
                  <Input
                    value={values.BooleanValueDefault}
                    name="BooleanValueDefault"
                    onChange={handleChange}
                    onBlur={handleBlur}
                  />
                </Widget>
                <Widget>
                  <Label>DateTimeValueDefault</Label>
                  <Input
                    value={values.DateTimeValueDefault}
                    name="DateTimeValueDefault"
                    onChange={handleChange}
                    onBlur={handleBlur}
                  />
                </Widget>
                <Widget>
                  <Label>ComponentType</Label>
                  <Input
                    value={values.ComponentType}
                    name="ComponentType"
                    onChange={handleChange}
                    onBlur={handleBlur}
                  />
                </Widget>

                <Button
                  label="Save"
                  block
                  variant="primary"
                  onClick={handleSubmit}
                />
              </Form>
            </Card>
          </Content>
        )
      }}
    </Formik>
  )
}

EditFieldType.propTypes = {
  id: PropTypes.string,
}

export default EditFieldType
