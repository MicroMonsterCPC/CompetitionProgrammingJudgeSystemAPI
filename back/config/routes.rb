Rails.application.routes.draw do
  # For details on the DSL available within this file, see http://guides.rubyonrails.org/routing.html
  namespace :api do
    namespace :questions do
      get '/', action: 'index'
    end
  end

  get 'auth/:provider/callback', to: 'sessions#callback'
  get '/logout', to: 'sessions#destroy', as: :logout
  get '/success', to: 'home#success', as: :success

  # post "/answers/create", to: "answers#create", as: :create_answer

  root 'home#top'
  resources :questions
  resources :answers, :only => [:index, :create]
end
