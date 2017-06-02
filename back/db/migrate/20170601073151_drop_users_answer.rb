class DropUsersAnswer < ActiveRecord::Migration[5.0]
  def change
    add_column :answers, :user_id, :integer
    add_column :questions, :user_id, :integer
    drop_table :user_answers
  end
end
