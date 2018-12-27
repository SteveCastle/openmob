'use strict';

/**
 * Campaign.js controller
 *
 * @description: A set of functions called "actions" for managing `Campaign`.
 */

module.exports = {

  /**
   * Retrieve campaign records.
   *
   * @return {Object|Array}
   */

  find: async (ctx) => {
    if (ctx.query._q) {
      return strapi.services.campaign.search(ctx.query);
    } else {
      return strapi.services.campaign.fetchAll(ctx.query);
    }
  },

  /**
   * Retrieve a campaign record.
   *
   * @return {Object}
   */

  findOne: async (ctx) => {
    return strapi.services.campaign.fetch(ctx.params);
  },

  /**
   * Count campaign records.
   *
   * @return {Number}
   */

  count: async (ctx) => {
    return strapi.services.campaign.count(ctx.query);
  },

  /**
   * Create a/an campaign record.
   *
   * @return {Object}
   */

  create: async (ctx) => {
    return strapi.services.campaign.add(ctx.request.body);
  },

  /**
   * Update a/an campaign record.
   *
   * @return {Object}
   */

  update: async (ctx, next) => {
    return strapi.services.campaign.edit(ctx.params, ctx.request.body) ;
  },

  /**
   * Destroy a/an campaign record.
   *
   * @return {Object}
   */

  destroy: async (ctx, next) => {
    return strapi.services.campaign.remove(ctx.params);
  },

  /**
   * Add relation to a/an campaign record.
   *
   * @return {Object}
   */

  createRelation: async (ctx, next) => {
    return strapi.services.campaign.addRelation(ctx.params, ctx.request.body);
  },

  /**
   * Update relation to a/an campaign record.
   *
   * @return {Object}
   */

  updateRelation: async (ctx, next) => {
    return strapi.services.campaign.editRelation(ctx.params, ctx.request.body);
  },

  /**
   * Destroy relation to a/an campaign record.
   *
   * @return {Object}
   */

  destroyRelation: async (ctx, next) => {
    return strapi.services.campaign.removeRelation(ctx.params, ctx.request.body);
  }
};
