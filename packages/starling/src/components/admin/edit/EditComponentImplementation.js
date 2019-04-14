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

const GET_COMPONENTIMPLEMENTATION = gql`
  query getComponentImplementationById($id: ID!) {
    getComponentImplementation(ID: $id) {
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
      Path
      ComponentType {
        ID
      }
    }
  }
`
const UPDATE_COMPONENTIMPLEMENTATION = gql`
  mutation updateComponentImplementation(
    $id: ID!
    $componentImplementation: ComponentImplementationInput
  ) {
    updateComponentImplementation(
      ID: $id
      componentImplementation: $componentImplementation
      buildStatic: true
    )
  }
`

function EditComponentImplementation({ id }) {
  const {
    data: { getComponentImplementation: item = {} },
    error,
    loading,
  } = useQuery(GET_COMPONENTIMPLEMENTATION, {
    variables: { id },
  })

  const updateComponentImplementation = useMutation(
    UPDATE_COMPONENTIMPLEMENTATION
  )

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
        Path: parseObject(item.Path),
        ComponentType: parseObject(item.ComponentType),
      }}
      onSubmit={(values, { setSubmitting }) =>
        updateComponentImplementation({
          variables: {
            id: item.ID,
            componentImplementation: {
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
                    type="text"
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
                    type="text"
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
                    type="text"
                    onChange={handleChange}
                    onBlur={handleBlur}
                  />
                </Widget>
                <Widget>
                  <Label>Title</Label>
                  <Input
                    value={values.Title}
                    name="Title"
                    type="text"
                    onChange={handleChange}
                    onBlur={handleBlur}
                  />
                </Widget>
                <Widget>
                  <Label>Path</Label>
                  <Input
                    value={values.Path}
                    name="Path"
                    type="text"
                    onChange={handleChange}
                    onBlur={handleBlur}
                  />
                </Widget>
                <Widget>
                  <Label>ComponentType</Label>
                  <Input
                    value={values.ComponentType}
                    name="ComponentType"
                    type="text"
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

EditComponentImplementation.propTypes = {
  id: PropTypes.string,
}

export default EditComponentImplementation
