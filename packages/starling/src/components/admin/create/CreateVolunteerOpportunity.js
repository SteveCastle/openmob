import React from 'react'
import PropTypes from 'prop-types'
import Input from '@openmob/bluebird/src/components/forms/Input'
import Button from '@openmob/bluebird/src/components/buttons/Button'

const CreateVolunteerOpportunity = ({ id }) => (
  <div>
    <h1>Create VolunteerOpportunity</h1>
    <div>
      <label>ID</label>
      <Input placeholder="ID!" />
    </div>
    <div>
      <label>CreatedAt</label>
      <Input placeholder="Time!" />
    </div>
    <div>
      <label>UpdatedAt</label>
      <Input placeholder="Time!" />
    </div>
    <div>
      <label>Title</label>
      <Input placeholder="String!" />
    </div>
    <div>
      <label>VolunteerOpportunityType</label>
      <Input placeholder="ID" />
    </div>

    <Button label="Edit" />
  </div>
)

CreateVolunteerOpportunity.propTypes = {
  id: PropTypes.string,
}

export default CreateVolunteerOpportunity
