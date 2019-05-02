/* eslint-disable */

import React from 'react'
import PropTypes from 'prop-types'
import { useMutation } from 'react-apollo-hooks'
import gql from 'graphql-tag'
import { Formik } from 'formik'
import Content from '@openmob/bluebird/src/components/layout/Content'
import Card from '@openmob/bluebird/src/components/cards/Card'
import Form from '@openmob/bluebird/src/components/forms/Form'
import Widget from '@openmob/bluebird/src/components/forms/Widget'
import Label from '@openmob/bluebird/src/components/forms/Label'
import Input from '@openmob/bluebird/src/components/forms/Input'
import TextArea from '@openmob/bluebird/src/components/forms/TextArea'
import Button from '@openmob/bluebird/src/components/buttons/Button'

const CREATE_LAYOUTROW = gql`
  mutation createLayoutRow($layoutRow: LayoutRowInput) {
    createLayoutRow(layoutRow: $layoutRow, buildStatic: true) {
      ID
    }
  }
`

const CreateLayoutRow = ({ id }) => {
  const createLayoutRow = useMutation(CREATE_LAYOUTROW)

  return (
    <Formik
      onSubmit={(values, { setSubmitting }) =>
        createLayoutRow({
          variables: {
            layoutRow: {
              ...values,
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
                <h1>Create LayoutRow</h1>
                <Widget>
                  <Label>Layout</Label>
                  <Input
                    value={values.Layout}
                    type="text"
                    name="Layout"
                    onChange={handleChange}
                    onBlur={handleBlur}
                  />
                </Widget>
                <Widget>
                  <Label>Container</Label>
                  <Input
                    value={values.Container}
                    type="checkbox"
                    name="Container"
                    onChange={handleChange}
                    onBlur={handleBlur}
                  />
                </Widget>
                <Widget>
                  <Label>Weight</Label>
                  <Input
                    value={values.Weight}
                    type="number"
                    name="Weight"
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

CreateLayoutRow.propTypes = {
  id: PropTypes.string,
}

export default CreateLayoutRow
