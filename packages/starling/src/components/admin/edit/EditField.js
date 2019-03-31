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

const GET_FIELD = gql`
  query getFieldById($id: ID!) {
    getField(ID: $id) {
      ID
      CreatedAt {
        seconds
        nanos
      }
      UpdatedAt {
        seconds
        nanos
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
        nanos
      }
      Component {
        ID
      }
    }
  }
`
const UPDATE_FIELD = gql`
  mutation updateField($id: ID!, $field: FieldInput) {
    updateField(ID: $id, field: $field, buildStatic: false)
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

  const updateField = useMutation(UPDATE_FIELD)

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
        FieldType: parseObject(item.FieldType),
        StringValue: parseObject(item.StringValue),
        IntValue: parseObject(item.IntValue),
        FloatValue: parseObject(item.FloatValue),
        BooleanValue: parseObject(item.BooleanValue),
        DateTimeValue: parseObject(item.DateTimeValue),
        Component: parseObject(item.Component),
      }}
      onSubmit={(values, { setSubmitting }) =>
        updateField({
          variables: {
            id: item.ID,
            field: {
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
                  <Label>FieldType</Label>
                  <Input
                    value={values.FieldType}
                    name="FieldType"
                    onChange={handleChange}
                    onBlur={handleBlur}
                  />
                </Widget>
                <Widget>
                  <Label>StringValue</Label>
                  <Input
                    value={values.StringValue}
                    name="StringValue"
                    onChange={handleChange}
                    onBlur={handleBlur}
                  />
                </Widget>
                <Widget>
                  <Label>IntValue</Label>
                  <Input
                    value={values.IntValue}
                    name="IntValue"
                    onChange={handleChange}
                    onBlur={handleBlur}
                  />
                </Widget>
                <Widget>
                  <Label>FloatValue</Label>
                  <Input
                    value={values.FloatValue}
                    name="FloatValue"
                    onChange={handleChange}
                    onBlur={handleBlur}
                  />
                </Widget>
                <Widget>
                  <Label>BooleanValue</Label>
                  <Input
                    value={values.BooleanValue}
                    name="BooleanValue"
                    onChange={handleChange}
                    onBlur={handleBlur}
                  />
                </Widget>
                <Widget>
                  <Label>DateTimeValue</Label>
                  <Input
                    value={values.DateTimeValue}
                    name="DateTimeValue"
                    onChange={handleChange}
                    onBlur={handleBlur}
                  />
                </Widget>
                <Widget>
                  <Label>Component</Label>
                  <Input
                    value={values.Component}
                    name="Component"
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

EditField.propTypes = {
  id: PropTypes.string,
}

export default EditField
