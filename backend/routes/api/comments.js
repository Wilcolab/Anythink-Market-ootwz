/**
 * Express router for handling comment-related API endpoints.
 * 
 * @module routes/api/comments
 */

 /**
    * GET /
    * Retrieves all comments from the database.
    *
    * @name GET/
    * @function
    * @memberof module:routes/api/comments
    * @async
    * @param {express.Request} req - Express request object
    * @param {express.Response} res - Express response object
    * @returns {Object[]} 200 - Array of comment objects
    * @returns {Object} 500 - Internal server error
    */

 /**
    * DELETE /:id
    * Deletes a comment by its ID.
    *
    * @name DELETE/:id
    * @function
    * @memberof module:routes/api/comments
    * @async
    * @param {express.Request} req - Express request object
    * @param {express.Response} res - Express response object
    * @returns {void} 204 - No content, successful deletion
    * @returns {Object} 500 - Internal server error
    */
const router = require("express").Router();
const mongoose = require("mongoose");
const Comment = mongoose.model("Comment");

module.exports = router;


router.get("/", async (req, res) => {
  try {
    const comments = await Comment.find();
    res.json(comments);
  } catch (error) {
    res.status(500).json({ error: "Internal server error" });
  }
});

router.delete("/:id", async (req, res) => {
  try {
    const { id } = req.params;
    await Comment.findByIdAndDelete(id);
    res.status(204).send();
  } catch (error) {
    res.status(500).json({ error: "Internal server error" });
  }
});
