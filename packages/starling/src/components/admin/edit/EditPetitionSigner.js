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

const GET_PETITIONSIGNER = gql`
  query getPetitionSignerById($id: ID!) {
    getPetitionSigner(ID: $id) {
      ID
      CreatedAt {
        seconds
        nanos
      }
      UpdatedAt {
        seconds
        nanos
      }
      Petition {
        ID
      }
      Contact {
        ID
      }
      Cause {
        ID
      }
    }
  }
`
const UPDATE_PETITIONSIGNER = gql`
  mutation updatePetitionSigner(
    $id: ID!
    $petitionSigner: PetitionSignerInput
  ) {
    updatePetitionSigner(
      ID: $id
      petitionSigner: $petitionSigner
      buildStatic: true
    )
  }
`

function EditPetitionSigner({ id }) {
  const {
    data: { getPetitionSigner: item = {} },
    error,
    loading,
  } = useQuery(GET_PETITIONSIGNER, {
    variables: { id },
  })

  const updatePetitionSigner = useMutation(UPDATE_PETITIONSIGNER)

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
        Petition: parseObject(item.Petition),
        Contact: parseObject(item.Contact),
        Cause: parseObject(item.Cause),
      }}
      onSubmit={(values, { setSubmitting }) =>
        updatePetitionSigner({
          variables: {
            id: item.ID,
            petitionSigner: {
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
                  <Label>Petition</Label>
                  <Input
                    value={values.Petition}
                    name="Petition"
                    type="text"
                    onChange={handleChange}
                    onBlur={handleBlur}
                  />
                </Widget>
                <Widget>
                  <Label>Contact</Label>
                  <Input
                    value={values.Contact}
                    name="Contact"
                    type="text"
                    onChange={handleChange}
                    onBlur={handleBlur}
                  />
                </Widget>
                <Widget>
                  <Label>Cause</Label>
                  <Input
                    value={values.Cause}
                    name="Cause"
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

EditPetitionSigner.propTypes = {
  id: PropTypes.string,
}

export default EditPetitionSigner
