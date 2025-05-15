/**
 * @route GET /
 * @group Comments - Operations about comments
 * @returns {Array.<Object>} 200 - An array of comment objects with user info
 * @returns {Error} 500 - Error fetching comments
 * @description Retrieves all comments from the database, populating the 'user' field with the username.
 */

/**
 * @route DELETE /:id
 * @group Comments - Operations about comments
 * @param {string} id.path.required - The ID of the comment to delete
 * @returns {Object} 200 - Success message
 * @returns {Error} 404 - Comment not found
 * @returns {Error} 500 - Error deleting comment
 * @description Deletes a comment by its ID.
 */
const router = require("express").Router();
const mongoose = require("mongoose");
const Comment = mongoose.model("Comment");

module.exports = router;

router.get("/", async (req, res) => {
  try {
    const comments = await Comment.find().populate("user", "username");
    res.status(200).json(comments);
  } catch (error) {
    res.status(500).json({ message: "Error fetching comments" });
  }
}
);


// add an encpoint to delete comments
router.delete("/:id", async (req, res) => {
  try {
    const comment = await Comment.findByIdAndDelete(req.params.id);
    if (!comment) {
      return res.status(404).json({ message: "Comment not found" });
    }
    res.status(200).json({ message: "Comment deleted successfully" });
  } catch (error) {
    res.status(500).json({ message: "Error deleting comment" });
  }
});

