import React from 'react'
import PropTypes from 'prop-types'
import Input from '@openmob/bluebird/src/components/forms/Input'
import Button from '@openmob/bluebird/src/components/buttons/Button'

const CreateContact = ({ id }) => (
  <div>
    <h1>Create Contact</h1>
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
      <label>FirstName</label>
      <Input placeholder="String" />
    </div>
    <div>
      <label>MiddleName</label>
      <Input placeholder="String" />
    </div>
    <div>
      <label>LastName</label>
      <Input placeholder="String" />
    </div>
    <div>
      <label>Email</label>
      <Input placeholder="String" />
    </div>
    <div>
      <label>PhoneNumber</label>
      <Input placeholder="String" />
    </div>

    <Button label="Edit" />
  </div>
)

CreateContact.propTypes = {
  id: PropTypes.string,
}

export default CreateContact
