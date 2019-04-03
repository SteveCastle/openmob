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

const GET_MAILINGADDRESS = gql`
  query getMailingAddressById($id: ID!) {
    getMailingAddress(ID: $id) {
      ID
      CreatedAt {
        seconds
        nanos
      }
      UpdatedAt {
        seconds
        nanos
      }
      StreetAddress
      City
      State
      ZipCode
    }
  }
`
const UPDATE_MAILINGADDRESS = gql`
  mutation updateMailingAddress(
    $id: ID!
    $mailingAddress: MailingAddressInput
  ) {
    updateMailingAddress(
      ID: $id
      mailingAddress: $mailingAddress
      buildStatic: true
    )
  }
`

function EditMailingAddress({ id }) {
  const {
    data: { getMailingAddress: item = {} },
    error,
    loading,
  } = useQuery(GET_MAILINGADDRESS, {
    variables: { id },
  })

  const updateMailingAddress = useMutation(UPDATE_MAILINGADDRESS)

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
        StreetAddress: parseObject(item.StreetAddress),
        City: parseObject(item.City),
        State: parseObject(item.State),
        ZipCode: parseObject(item.ZipCode),
      }}
      onSubmit={(values, { setSubmitting }) =>
        updateMailingAddress({
          variables: {
            id: item.ID,
            mailingAddress: {
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
                  <Label>StreetAddress</Label>
                  <Input
                    value={values.StreetAddress}
                    name="StreetAddress"
                    onChange={handleChange}
                    onBlur={handleBlur}
                  />
                </Widget>
                <Widget>
                  <Label>City</Label>
                  <Input
                    value={values.City}
                    name="City"
                    onChange={handleChange}
                    onBlur={handleBlur}
                  />
                </Widget>
                <Widget>
                  <Label>State</Label>
                  <Input
                    value={values.State}
                    name="State"
                    onChange={handleChange}
                    onBlur={handleBlur}
                  />
                </Widget>
                <Widget>
                  <Label>ZipCode</Label>
                  <Input
                    value={values.ZipCode}
                    name="ZipCode"
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

EditMailingAddress.propTypes = {
  id: PropTypes.string,
}

export default EditMailingAddress
