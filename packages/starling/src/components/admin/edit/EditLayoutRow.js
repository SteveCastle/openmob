/* eslint-disable */

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
import TextArea from '@openmob/bluebird/src/components/forms/TextArea'
import Button from '@openmob/bluebird/src/components/buttons/Button'
import parseObject from '../../../common/helpers'

const GET_LAYOUTROW = gql`
  query getLayoutRowById($id: ID!) {
    getLayoutRow(ID: $id) {
      ID
      CreatedAt {
        seconds
        nanos
      }
      UpdatedAt {
        seconds
        nanos
      }
      Layout {
        ID
      }
      Container
      Weight
    }
  }
`
const UPDATE_LAYOUTROW = gql`
  mutation updateLayoutRow($id: ID!, $layoutRow: LayoutRowInput) {
    updateLayoutRow(ID: $id, layoutRow: $layoutRow, buildStatic: true)
  }
`

function EditLayoutRow({ id }) {
  const {
    data: { getLayoutRow: item = {} },
    error,
    loading,
  } = useQuery(GET_LAYOUTROW, {
    variables: { id },
  })

  const updateLayoutRow = useMutation(UPDATE_LAYOUTROW)

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
        Layout: parseObject(item.Layout),
        Container: parseObject(item.Container),
        Weight: parseObject(item.Weight),
      }}
      onSubmit={(values, { setSubmitting }) =>
        updateLayoutRow({
          variables: {
            id: item.ID,
            layoutRow: {
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
                  <Label>Layout</Label>
                  <Input
                    value={values.Layout}
                    name="Layout"
                    type="text"
                    onChange={handleChange}
                    onBlur={handleBlur}
                  />
                </Widget>
                <Widget>
                  <Label>Container</Label>
                  <Input
                    value={values.Container}
                    name="Container"
                    type="checkbox"
                    onChange={handleChange}
                    onBlur={handleBlur}
                  />
                </Widget>
                <Widget>
                  <Label>Weight</Label>
                  <Input
                    value={values.Weight}
                    name="Weight"
                    type="number"
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

EditLayoutRow.propTypes = {
  id: PropTypes.string,
}

export default EditLayoutRow
