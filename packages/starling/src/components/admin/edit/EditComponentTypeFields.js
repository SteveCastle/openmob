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
import parseObject from '../../../common/helpers'

const GET_COMPONENTTYPEFIELDS = gql`
  query getComponentTypeFieldsById($id: ID!) {
    getComponentTypeFields(ID: $id) {
      ID
      CreatedAt {
        seconds
        nanos
      }
      UpdatedAt {
        seconds
        nanos
      }
      ComponentType {
        ID
      }
      FieldType {
        ID
      }
      Weight
      Required
    }
  }
`
const UPDATE_COMPONENTTYPEFIELDS = gql`
  mutation updateComponentTypeFields(
    $id: ID!
    $componentTypeFields: ComponentTypeFieldsInput
  ) {
    updateComponentTypeFields(
      ID: $id
      componentTypeFields: $componentTypeFields
      buildStatic: true
    )
  }
`

function EditComponentTypeFields({ id }) {
  const {
    data: { getComponentTypeFields: item = {} },
    error,
    loading,
  } = useQuery(GET_COMPONENTTYPEFIELDS, {
    variables: { id },
  })

  const updateComponentTypeFields = useMutation(UPDATE_COMPONENTTYPEFIELDS)

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
        ComponentType: parseObject(item.ComponentType),
        FieldType: parseObject(item.FieldType),
        Weight: parseObject(item.Weight),
        Required: parseObject(item.Required),
      }}
      onSubmit={(values, { setSubmitting }) =>
        updateComponentTypeFields({
          variables: {
            id: item.ID,
            componentTypeFields: {
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
                  <Label>ComponentType</Label>
                  <Input
                    value={values.ComponentType}
                    name="ComponentType"
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
                  <Label>Weight</Label>
                  <Input
                    value={values.Weight}
                    name="Weight"
                    onChange={handleChange}
                    onBlur={handleBlur}
                  />
                </Widget>
                <Widget>
                  <Label>Required</Label>
                  <Input
                    value={values.Required}
                    name="Required"
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

EditComponentTypeFields.propTypes = {
  id: PropTypes.string,
}

export default EditComponentTypeFields
