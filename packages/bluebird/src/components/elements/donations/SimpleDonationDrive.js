import React from 'react';
import styled from 'styled-components';
import PropTypes from 'prop-types';

import {
  space,
  width,
  fontSize,
  color,
  borders,
  borderRadius,
  display,
  alignItems,
  justifyContent,
} from 'styled-system';

const StyledDonationDrive = styled.section`
  ${borders}
  ${borderRadius}
  ${space}
  ${width}
  ${fontSize}
  ${alignItems}
  ${justifyContent}
  ${color}
  ${display}
  min-height: 400px;
  `;
const Title = styled.h1`
  ${borders}
  ${borderRadius}
  ${space}
  ${width}
  ${fontSize}
  ${alignItems}
  ${justifyContent}
  ${color}
  ${display}
  `;

function SimpleDonationDrive({ title }) {
  return (
    <StyledDonationDrive
      display={'flex'}
      justifyContent="center"
      alignItems="center"
      bg="lime.4"
    >
      <Title color="gray.1" size={1}>
        {title}
      </Title>
    </StyledDonationDrive>
  );
}

SimpleDonationDrive.propTypes = {
  title: PropTypes.string.isRequired,
};
SimpleDonationDrive.defaultProps = { title: 'Simple Donation Drive' };

export default SimpleDonationDrive;
