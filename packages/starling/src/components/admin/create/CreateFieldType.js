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
import Button from '@openmob/bluebird/src/components/buttons/Button'

const CREATE_FIELDTYPE = gql`
  mutation createFieldType($fieldType: FieldTypeInput) {
    createFieldType(fieldType: $fieldType, buildStatic: true) {
      ID
    }
  }
`

const CreateFieldType = ({ id }) => {
  const createFieldType = useMutation(CREATE_FIELDTYPE)

  return (
    <Formik
      onSubmit={(values, { setSubmitting }) =>
        createFieldType({
          variables: {
            fieldType: {
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
                <h1>Create FieldType</h1>
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

CreateFieldType.propTypes = {
  id: PropTypes.string,
}

export default CreateFieldType
