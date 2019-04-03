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

const GET_VOLUNTEEROPPORTUNITY = gql`
  query getVolunteerOpportunityById($id: ID!) {
    getVolunteerOpportunity(ID: $id) {
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
      VolunteerOpportunityType {
        ID
      }
    }
  }
`
const UPDATE_VOLUNTEEROPPORTUNITY = gql`
  mutation updateVolunteerOpportunity(
    $id: ID!
    $volunteerOpportunity: VolunteerOpportunityInput
  ) {
    updateVolunteerOpportunity(
      ID: $id
      volunteerOpportunity: $volunteerOpportunity
      buildStatic: true
    )
  }
`

function EditVolunteerOpportunity({ id }) {
  const {
    data: { getVolunteerOpportunity: item = {} },
    error,
    loading,
  } = useQuery(GET_VOLUNTEEROPPORTUNITY, {
    variables: { id },
  })

  const updateVolunteerOpportunity = useMutation(UPDATE_VOLUNTEEROPPORTUNITY)

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
        VolunteerOpportunityType: parseObject(item.VolunteerOpportunityType),
      }}
      onSubmit={(values, { setSubmitting }) =>
        updateVolunteerOpportunity({
          variables: {
            id: item.ID,
            volunteerOpportunity: {
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
                  <Label>Title</Label>
                  <Input
                    value={values.Title}
                    name="Title"
                    onChange={handleChange}
                    onBlur={handleBlur}
                  />
                </Widget>
                <Widget>
                  <Label>VolunteerOpportunityType</Label>
                  <Input
                    value={values.VolunteerOpportunityType}
                    name="VolunteerOpportunityType"
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

EditVolunteerOpportunity.propTypes = {
  id: PropTypes.string,
}

export default EditVolunteerOpportunity
