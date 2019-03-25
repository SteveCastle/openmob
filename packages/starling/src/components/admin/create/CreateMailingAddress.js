import React from 'react'
import PropTypes from 'prop-types'
import Input from '@openmob/bluebird/src/components/forms/Input'
import Button from '@openmob/bluebird/src/components/buttons/Button'

const CreateMailingAddress = ({ id }) => (
  <div>
    <h1>Create MailingAddress</h1>
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
      <label>StreetAddress</label>
      <Input placeholder="String!" />
    </div>
    <div>
      <label>City</label>
      <Input placeholder="String!" />
    </div>
    <div>
      <label>State</label>
      <Input placeholder="String!" />
    </div>
    <div>
      <label>ZipCode</label>
      <Input placeholder="String!" />
    </div>

    <Button label="Edit" />
  </div>
)

CreateMailingAddress.propTypes = {
  id: PropTypes.string,
}

export default CreateMailingAddress
