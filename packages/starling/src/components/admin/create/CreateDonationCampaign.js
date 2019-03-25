import React from 'react'
import PropTypes from 'prop-types'
import Input from '@openmob/bluebird/src/components/forms/Input'
import Button from '@openmob/bluebird/src/components/buttons/Button'

const CreateDonationCampaign = ({ id }) => (
  <div>
    <h1>Create DonationCampaign</h1>
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

    <Button label="Edit" />
  </div>
)

CreateDonationCampaign.propTypes = {
  id: PropTypes.string,
}

export default CreateDonationCampaign
