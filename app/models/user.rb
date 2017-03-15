class User < ApplicationRecord

  has_many :user_answers
  has_many :users, :through => :user_answers

  def self.create_with_omniauth(auth)
    create! do |user|
      user.provider = auth['provider']
      user.uid = auth['uid']
      user.name = auth['info']['nickname']
      user.oauth_token = auth['credentials']['token']
    end
  end
end
