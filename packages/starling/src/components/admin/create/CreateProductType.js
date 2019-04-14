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

const CREATE_PRODUCTTYPE = gql`
  mutation createProductType($productType: ProductTypeInput) {
    createProductType(productType: $productType, buildStatic: true) {
      ID
    }
  }
`

const CreateProductType = ({ id }) => {
  const createProductType = useMutation(CREATE_PRODUCTTYPE)

  return (
    <Formik
      onSubmit={(values, { setSubmitting }) =>
        createProductType({
          variables: {
            productType: {
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
                <h1>Create ProductType</h1>
                <Widget>
                  <Label>Title</Label>
                  <Input
                    value={values.Title}
                    type="text"
                    name="Title"
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

CreateProductType.propTypes = {
  id: PropTypes.string,
}

export default CreateProductType
