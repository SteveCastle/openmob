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

const CREATE_LAYOUTCOLUMN = gql`
  mutation createLayoutColumn($layoutColumn: LayoutColumnInput) {
    createLayoutColumn(layoutColumn: $layoutColumn, buildStatic: true) {
      ID
    }
  }
`

const CreateLayoutColumn = ({ id }) => {
  const createLayoutColumn = useMutation(CREATE_LAYOUTCOLUMN)

  return (
    <Formik
      onSubmit={(values, { setSubmitting }) =>
        createLayoutColumn({
          variables: {
            layoutColumn: {
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
                <h1>Create LayoutColumn</h1>
                <Widget>
                  <Label>LayoutRow</Label>
                  <Input
                    value={values.LayoutRow}
                    type="text"
                    name="LayoutRow"
                    onChange={handleChange}
                    onBlur={handleBlur}
                  />
                </Widget>
                <Widget>
                  <Label>Width</Label>
                  <Input
                    value={values.Width}
                    type="number"
                    name="Width"
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

CreateLayoutColumn.propTypes = {
  id: PropTypes.string,
}

export default CreateLayoutColumn
