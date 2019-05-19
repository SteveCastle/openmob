import React from 'react';
import { shallow } from 'enzyme';
import SubHeader from '../SubHeader';

it('renders without crashing', () => {
  shallow(<SubHeader>Cool title.</SubHeader>);
});
