class HomeController < ApplicationController
  def top
    render 'home/top.html.erb'
  end
  def success
    render 'home/success.html.erb'
  end
end
